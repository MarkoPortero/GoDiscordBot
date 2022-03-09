package GlobalVariables

import "os"

func LoadEnvVariables() {
	MongoUri = os.Getenv("MONGO_URI")
}
