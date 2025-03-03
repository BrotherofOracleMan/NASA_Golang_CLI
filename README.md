# NASA_Golang_CLI
Golang Project for teaching me Golang Concepts

See NASA's api at https://api.nasa.gov/.

Functionality

- Fetch Astronomy Picture of the Day
  - Create flags for date and put the command into a history.
  - Do a Validation of the date before fetching. If the date is not given, set it to the current day.
  - Put command name image url into a database
   
- Fetch Earth Imagery data  
  - Create flags for lat, long, dim and date
  - Do a Validation of the above parameters mentioned. Make sure to include default parameters
  - Include a parameter that downloads assets. Maybe use goroutines and workers to download mulitple images.
  - Put command name/ image url into a data base.

- Fetch Mars Rover Photos
  - TODO Add:

- Fetch the most Recent X Commands. Display the Images or Image storage url if there are alot of images.
