
## Required Helm repos for NGINX Ingress Controller and web-app deployment:

**Kubernetes - Google**
```helm repo add stable https://kubernetes-charts.storage.googleapis.com/```

**NGINX Ingress Controller** *(taken from this repo, slightly customized and saved locally in git to be deployed with the customization included)*
```helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx```

**Cert-Manager**
```helm repo add jetstack https://charts.jetstack.io```

## Docker Hub Repository
My repository that stores the containers used by K8s can be found here:
```https://hub.docker.com/u/nirbo2```

## Web Application
The web-app is available via: https://newcomers.internal.fin.ag


## REST API Documentation

The REST API is accessible via https://newcomers-api.internal.fin.ag/persons
It is built around the "Person" struct and each object contains a person's details, as shown in the example below.
The API is capable of five basic CRUD operations, listed and detailed below as well, and includes an additional /health endpoint for the DaemonSet's readinessProbe.

**This API has no persistence and is completely dependent on the container's memory state.**

#### JSON object structure example:
```
[
	{
	"first_name": "Nir",
	"last_name": "Ben-Or",
	"physical_attributes": {
		"height": "178cm",
		"eye_color": "Brown"
	}
]
```

#### Functions:
1.	**getAllPersons()**: Retrieve all Person objects.

	Allowed methods: **GET**

	**Usage:**
	```bash
	curl -X GET https://newcomers-api.internal.fin.ag/persons 
	```
<br>

2.	**getPersonByID()**: Retrieve one Person object using its JSON array index ID.

	Allowed methods: **GET**

	**Usage:**
	```bash
	curl -X GET https://newcomers-api.internal.fin.ag/persons/{0..N} 
	```
<br>

3.	**addPerson()**: Add a new Person object.

	Allowed methods: **POST**
	
	**Usage:**
	```bash
	 curl -X POST  -d ' \
        {
            "first_name": "Nir",
            "last_name": "Ben-Or",
            "physical_attributes": {
                "height": "178cm",
                "eye_color": "Brown"
            }
        }' \
        https://newcomers-api.internal.fin.ag/persons
	```
<br>
	
4.	**updatePersonByID()**: Update one or more fields of a single Person object that already exists in the array using its JSON array index ID.

	*Note: The current implementation requires all fields to be populated, including the fields that are not being updated.*
	
	Allowed methods: **PUT**
	
	**Usage:**
	```bash
	 curl -X PUT  -d ' \
        {
            "first_name": "Nir",
            "last_name": "Ben-Or",
            "physical_attributes": {
                "height": "160cm",
                "eye_color": "Blue"
            }
        }' \
        https://newcomers-api.internal.fin.ag/persons/{0..N}
	```
<br>

5.	**deletePersonByID()**: Delete one Person object using its JSON array index ID.

	Allowed methods: **DELETE**

	**Usage:**
	```bash
	curl -X DELETE https://newcomers-api.internal.fin.ag/persons/{0..N} 
	```
<br>

6.	**healthCheck()**: A method accessible by https://newcomers-api.internal.fin.ag/health that is used for checking the health of the REST API application.
        This method is used only by the REST API's K8s Deployment 'readinessProbe'.
	It always returns HTTP Code 200.

	Allowed methods: **GET**
