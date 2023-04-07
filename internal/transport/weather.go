package transport

var WeatherData *CurrentWeatherData

type (
	CurrentWeatherData struct {
		GeoPos Coordinates
		Main   Main
	}
	Coordinates struct {
		Longitude float64
		Latitude  float64
		Location  string
	}
	Main struct {
		Temp      float64
		TempMin   float64
		TempMax   float64
		FeelsLike float64
		Pressure  float64
		Humidity  int
	}
)

func (c *CurrentWeatherData) GetWeatherParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) {
	return c.GeoPos.Longitude, c.GeoPos.Latitude, c.Main.Temp, c.Main.TempMin, c.Main.TempMax, c.Main.FeelsLike, c.Main.Pressure, c.Main.Humidity
}

var (
	APIKey      KeyAPIGetter
	Key, Answer string
)

type (
	API struct {
		Key string
	}
	KeyAPIGetter interface {
		getKey() string
	}
)

func (a *API) getKey() string {
	return a.Key
}

func GetAPIKey(key KeyAPIGetter) string {
	return key.getKey()
}
