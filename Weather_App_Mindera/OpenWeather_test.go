package Weather_App_Mindera

import(

	testing2 "testing"
)

func BenchmarkGetOpenWeather(b *testing2.B) {

   for i:=0;i<=b.N;i++{
	   GetOpenWeather()
   }

}
