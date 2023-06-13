![GitHub](https://img.shields.io/github/license/bradmccoydev/surf-report)
# Surf Report
Surf Report API. This used to be in donetcore for Microsoft Reactor labs ran by Brad, now it has been refactored to use Golang. For references to the labs you can checkout an earlier commit for the dotnet core code.

This application simply calls the weather api and creates prometheus metrics for use of viewing data related to surfing in tools like grafana.

curl "https://api.stormglass.io/v2/weather/point?lat=-33.7979&lng=151.2882&params=airTemperature,cloudCover,swellDirection,swellHeight,swellPeriod,waterTemperature,waveDirection,waveHeight" -H "Authorization: " > surf.json

curl "https://api.stormglass.io/v2/tide/sea-level/point?lat=-33.7979&lng=151.2882"  -H "Authorization: " > tide.json

# HELP surf_water_temperature Swell direction metric
# TYPE surf_water_temperature gauge
surf_water_temperature 18.42

# HELP surf_wave_direction Swell direction metric
# TYPE surf_wave_direction gauge
surf_wave_direction 152.14

# HELP surf_wave_height Swell height metric
# TYPE surf_wave_height gauge
surf_wave_height 152.14

# HELP surf_air_temperature Swell direction metric
# TYPE surf_air_temperature gauge
surf_air_temperature 16.45

# HELP surf_cloud_cover Swell direction metric
# TYPE surf_cloud_cover gauge
surf_cloud_cover 100.0

# HELP surf_swell_direction Swell direction metric
# TYPE surf_swell_direction gauge
surf_swell_direction 150.68

# HELP surf_swell_height Swell height metric
# TYPE surf_swell_height gauge
surf_swell_height 1.07

# HELP surf_swell_period Swell period metric
# TYPE surf_swell_period gauge
surf_swell_period 7.88

# HELP surf_tide Tide metric
# TYPE surf_tide gauge
surf_tide 7.88
