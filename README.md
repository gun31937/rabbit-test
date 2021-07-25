
# Rabbit Test

URL-shortener service to shorten URLs.
## Run Locally

Clone the project

```bash
  git clone https://github.com/gun31937/rabbit-test.git
```

Go to the project directory

```bash
  cd rabbit-test
```

Install dependencies

```bash
  go mod tidy
```

Start the server with Makefile

```bash
  make local
```

and server will start with port 8080
## Running Tests

To run tests, run the following command

```bash
  make test
```
## API Reference

### Create short url

```http
  POST /
```
#### Request
| Parameter   | Type     | Description                             |
| :---------- | :------- | :-------------------------------------- |
| `fullURL`   | `string` | **Required**. url you want to shorten   |
| `expiresIn` | `int`    | Created link will expires in (minutes). |

#### Response data
| Parameter  | Type     | Description                      |
| :--------- | :------- | :------------------------------- |
| `shortURL` | `string` | Shorten url.                     |
| `expiry`   | `string` | Date and time to expired (GMT+7) |

### Get item

```http
  GET /${shortCode}
```
#### Request
| Parameter    | Type     | Description                       |
| :----------- | :------- | :-------------------------------- |
| `shortCode`  | `string` | Short code (after base url)       |

#### Response
If url is valid, will redirect to full url

### Delete short url

```http
  DELETE /${shortCode}
```
#### Header
| Parameter       | Type     | Description                       |
| :-------------- | :------- | :-------------------------------- |
| `Authorization` | `string` | Bearer token                      |

#### Request
| Parameter    | Type     | Description                       |
| :----------- | :------- | :-------------------------------- |
| `shortCode`  | `string` | Short code (after base url)       |

#### Response
If no error, will return null data

### List url

```http
  GET /
```
#### Header
| Parameter       | Type     | Description                       |
| :-------------- | :------- | :-------------------------------- |
| `Authorization` | `string` | Bearer token                      |

#### Request
| Parameter    | Type     | Description                       |
| :----------- | :------- | :-------------------------------- |
| `shortCode`  | `string` | Short code (after base url)       |
| `keyword`    | `string` | Search on origin url              |

#### Response data array
| Parameter   | Type     | Description                      |
| :---------- | :------- | :------------------------------- |
| `id`        | `int`    | URL ID                           |
| `shortCode` | `string` | Shorten url.                     |
| `fullURL`   | `string` | Origin url.                      |
| `hits`      | `int`    | Number of hits                   |
| `expiry`    | `string` | Date and time to expired (GMT+7) |
| `createdAt` | `string` | Created date and time (GMT+7)    |
| `updatedAt` | `string` | Updated date and time (GMT+7)    |
| `deletedAt` | `string` | Deleted date and time (GMT+7)    |

### Login

```http
  POST /login
```

#### Request
| Parameter    | Type     | Description                       |
| :----------- | :------- | :-------------------------------- |
| `username`   | `string` | Use `admin` for demo              |
| `password`   | `string` | Use `admin` for demo              |

#### Response
| Parameter   | Type     | Description                      |
| :---------- | :------- | :------------------------------- |
| `token`     | `string` | Bearer token                     |
| `expire`    | `string` | Expiry                           |

and add bearer token to `Authorization` header to access admin endpoints