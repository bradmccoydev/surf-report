using System;
using System.Net;
using Newtonsoft.Json;
using System.Net.Http;
using System.Net.Http.Headers;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace SurfReport.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class WeatherForecastController : ControllerBase
    {
        private readonly ILogger<WeatherForecastController> _logger;

        public WeatherForecastController(ILogger<WeatherForecastController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public async Task<IActionResult> Get()
        {
            try
            {
                var apiToken = Environment.GetEnvironmentVariable("API_TOKEN");
                var spotId = Environment.GetEnvironmentVariable("SPOT_ID");

                if (apiToken == null || spotId == null)
                {
                    return Ok("No Token provided");
                }

                var client = GetHttpClient(mediaType: "application/json");
                string url = $"http://magicseaweed.com/api/{apiToken}/forecast?spot_id={spotId}";
                var stringTask = client.GetStringAsync(url);
                var json = await stringTask;
                var forecast = JsonConvert.DeserializeObject<List<WeatherForecast.RootObject>>(json);
                return Ok(json);
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "Surf Forecast error");
                return new StatusCodeResult((int)HttpStatusCode.InternalServerError);
            }

        }

        public HttpClient GetHttpClient(
            string mediaType)
        {
            var httpClient = new HttpClient();
            httpClient.DefaultRequestHeaders.Accept.Clear();
            httpClient.DefaultRequestHeaders.Accept.Add(
                new MediaTypeWithQualityHeaderValue(mediaType));
            TimeSpan ts = TimeSpan.FromSeconds(1000);
            httpClient.Timeout = ts;

            return httpClient;
        }
    }
}
