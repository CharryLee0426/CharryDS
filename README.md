# Charry Data Structure

## 0. Overview
Charry Data Structure is a repository that stores some useful data structures.
I use go programming language to implement these data structures due to three
advantages it has. First, go is a compiled language, which means it is faster
than many interpreted languages such as python. Second, it is very easy for
developers to design tests for their code. I love go test! Third, go is a
very widely used language right now, especially when it comes to cloud.

## 1. Data Structures
* Heap (min heap)
* Binary Tree (assumed to create a complete binary tree)
* Disjoint Set Data Structure (Union Find)

## 2. Algorithms
* Sort
    - [x] Heap Sort
    - [x] Insertion Sort
    - [x] Selection Sort
    - [x] Bubble Sort
    - [x] Quick Sort
    - [x] Merge Sort
    - [x] Pattern-defeating Quicksort (Pdqsort) (state of the art)

* Graph
    - [x] Kruskal
    - [x] Prim
    - [x] Dijkstra
    - [x] Bellman-Ford
    - [x] SPFA
    - [x] Kahn's algorithm (topological sort)

* Binary Search
    - [x] Template 1
        
        Why do we use `mid = left + (right - left) / 2` instead of `mid = (left + right) / 2`?

        Because the latter one may cause overflow when `left` and `right` are both very large.
        For example, if our maximum number is 100, left is 50, and right is 80, `left + right`
        is 130, which is larger than 100. Thus, it will cause an overflow. However, `right - left`
        is 30, which is smaller than 100. Thus, the former one doesn't cause an overflow.

Test condition:
Three types of data: shuffle, sorted, reversed, mod 8 (1k elements)
Disable compiler optimization
100 times for each type of data
Show memory allocation and number of allocations

Test result:
```
goos: darwin
goarch: arm64
pkg: charryds/Sort
BenchmarkBubbleSort-8                       1000           2467353 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-8                    1000           1227912 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-8                    1000           2048261 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeap-8                          1000           1024387 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeap-8                          1000            102687 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-8                        1000             55085 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-8                        1000             82767 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSort-8                          1000             47935 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsort-8                  1000             33705 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortSorted-8                 1000           1166343 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortSorted-8              1000           1177372 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortSorted-8              1000              3175 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapSorted-8                    1000            111988 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapSorted-8                    1000            107079 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortSorted-8                  1000           1146282 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortSorted-8                  1000             83550 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortSorted-8                    1000              2652 ns/op               0 B/op          0 allocs/op
BenchmarkPdqsortSorted-8                    1000              2603 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortReversed-8               1000           2980081 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortReversed-8            1000           1199145 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortReversed-8            1000           2914308 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapReversed-8                  1000           1032363 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapReversed-8                  1000             99003 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortReversed-8                1000           1318531 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortReversed-8                1000             82179 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortReversed-8                  1000              4139 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortReversed-8          1000              3852 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortMod8-8                   1000           1210913 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortMod8-8                1000           1189246 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortMod8-8                1000           1281706 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapMod8-8                      1000            707426 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapMod8-8                      1000             96785 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortMod8-8                    1000            194367 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortMod8-8                    1000             88986 ns/op           81920 B/op       1023 allocs/op
BenchmarkPDQSortMod8-8                      1000             18608 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortMod8-8              1000             13765 ns/op              24 B/op          1 allocs/op
BenchmarkBubbleSortRandom-8                 1000           2683426 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSortRandom-8              1000           1227405 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSortRandom-8              1000           1468105 ns/op               0 B/op          0 allocs/op
BenchmarkMinHeapRandom-8                    1000            748061 ns/op            9520 B/op          2 allocs/op
BenchmarkMaxHeapRandom-8                    1000            115578 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSortRandom-8                  1000             73461 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSortRandom-8                  1000            118175 ns/op           81922 B/op       1023 allocs/op
BenchmarkPDQSortRandom-8                    1000             82756 ns/op               0 B/op          0 allocs/op
BenchmarkOfficialPDQsortRandom-8            1000             63176 ns/op              24 B/op          1 allocs/op
```

* Recursive solution for eight queens puzzle

## 3. Backend Framework

`gin-gonic` is used widely as a backend framework in Golang. It is a high performance framework with a lot of features. Here is example code for using it.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin

func main() {
    // gin.DisableConsoleColor()
    // gin.ForceConsoleColor()

	// r := gin.Default()

    f, _ := os.Create("gin.log")

    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    r := gin.Default()

    // r := gin.New()

    // r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
    //     return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 			param.ClientIP,
	// 			param.TimeStamp.Format(time.RFC1123),
	// 			param.Method,
	// 			param.Path,
	// 			param.Request.Proto,
	// 			param.StatusCode,
	// 			param.Latency,
	// 			param.Request.UserAgent(),
	// 			param.ErrorMessage,
	// 	)
    // }))

    // r.Use(gin.Recovery())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/asciijson", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		ctx.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.LoadHTMLFiles("form.html", "index.html")
	// r.GET("/", indexHandler)
	// r.POST("/", formHandler)

    r.GET("/testing", startPage)

    r.GET("/person/:name/:address", func(ctx *gin.Context) {
        var person Person

        if err := ctx.ShouldBindUri(&person); err != nil {
            ctx.JSON(400, gin.H{"msg": err})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{"name": person.Name, "address": person.Address})
    })

    r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

    r.GET("/welcome", func(c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })

    // r.GET("/test", func(c *gin.Context) {
    //     c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
    // })

    r.GET("/test", func(c *gin.Context) {
        c.Request.URL.Path = "/test2"
        r.HandleContext(c)
    })

    r.GET("/test2", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"hello": "world"})
    })

    authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

    r.GET("/", func(ctx *gin.Context) {
        ctx.HTML(http.StatusOK, "index.html", nil)
    })

    r.POST("/upload", func(ctx *gin.Context) {
        name := ctx.PostForm("name")
        email := ctx.PostForm("email")

        form, err := ctx.MultipartForm()

        if err != nil {
            ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
        }

        files := form.File["files"]

        for _, file := range files {
            filename := filepath.Base(file.Filename)

            if err := ctx.SaveUploadedFile(file, "files/"+filename); err != nil {
                ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
            }
        }

        ctx.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
    })

    r.POST("/download", func(ctx *gin.Context) {
        form, err := ctx.MultipartForm()
        file := form.File["file"]

        filename := filepath.Base(file[0].Filename)

        if err != nil {
            ctx.String(http.StatusBadRequest, fmt.Sprintf("get file path err: %s", err.Error()))
        }

        ctx.FileAttachment("files/"+filename, filename)
        ctx.String(http.StatusOK, fmt.Sprintf("Download successfully %s", filename))
    })

    s := &http.Server{
        Addr:           ":8080",
        Handler:        r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}

type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(200, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}

type myForm struct {
    Colors []string `form:"colors[]"`
}

func indexHandler(c *gin.Context) {
    c.HTML(200, "form.html", nil)
}

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.Bind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}

type Person struct {
    Name string `form:"name" json:"name" uri:"name"`
    Address string `form:"address" json:"address" uri:"address"`
}

func startPage(ctx *gin.Context) {
    var person Person

    if ctx.Bind(&person) == nil {
        log.Println("====== Only Bind By Query String ======")
        log.Println(person.Name)
        log.Println(person.Address)
    }

    if ctx.BindJSON(&person) == nil {
        log.Println("====== Only Bind By JSON ======")
        log.Println(person.Name)
        log.Println(person.Address)
    }

    ctx.String(http.StatusOK, "Success")
}

// curl -X GET "localhost:8080/testing?name=appleboy&address=xyz"
// curl -X GET localhost:8080/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"

// BasicAuth middleware
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}
```

## 4. How to use jwt middleware

```go
var testusers = []User{
    User{ID: "1", Name: "Chen Li"},
    User{ID: "2", Name: "Chen Lii"},
    User{ID: "3", Name: "Chen Liii"},
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery())

    public := r.Group("/")
    private := r.Group("/")

    private.Use(JwtAuthMiddleware("secret"))

    public.POST("/login", func(c *gin.Context) {
        id := c.PostForm("id")

        for _, user := range testusers {
            if user.ID == id {
                token, err := CreateAccessToken(&user, "secret", 2)

                if err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                    c.Abort()
                }

                c.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
                c.JSON(http.StatusOK, gin.H{"token": token})
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{"error": "User not found"})
    })

    private.GET("/ping", func(c *gin.Context) {
        cookietoken, _ := c.Cookie("access_token")
        c.JSON(http.StatusOK, gin.H{"cookie": cookietoken})
    })

    r.Run(":8080")
}

type User struct {
    ID string
    Name string
}

type JwtCustomClaims struct {
    Name string `json:"name"`
    ID string `json:"id"`
    jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
    ID string `json:"id"`
    jwt.StandardClaims
}

func CreateAccessToken(user *User, secrect string, expiry int) (accessToken string, err error) {
    exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

    claims := &JwtCustomClaims{
        Name: user.Name,
        ID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: exp,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte(secrect))

    if err != nil {
        return "", err
    }

    return t, err
}

func CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error) {
    claimsRefresh := &JwtCustomRefreshClaims{
        ID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

    rt, err := token.SignedString([]byte(secret))

    if err != nil {
        return "", err
    }

    return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
    _, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
    token, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
        }

        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    claims, ok := token.Claims.(jwt.MapClaims)

    if !ok && !token.Valid {
        return "", fmt.Errorf("Invalid Token")
    }

    return claims["id"].(string), nil
}

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
        t := strings.Split(authHeader, " ")

        if len(t) == 2 {
            authToken := t[1]
            
            authorized, err := IsAuthorized(authToken, secret)

            if authorized {
                userID, err := ExtractIDFromToken(authToken, secret)

                if err != nil {
                    c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                    c.Abort()
                    return
                }

                c.Header("x-user-id", userID)
                c.Next()
                return
            }

            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
        }

        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
    }
}
```

## 5. How to use mysql driver

```go
type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User: "root",
		Passwd: "Lc15503539901",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",	
	}

	var err error

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	// ping
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	albums, err := fetchAlbumsByArtist("John Coltrane", db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	rows, err := db.Query("select * from album where artist = ?", "John Coltrane")

	albs := []Album{}

	for rows.Next() {
		var alb Album

		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

		if err != nil {
			log.Fatal(err)
		}

		albs = append(albs, alb)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", albums)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	newAlbID, err := updateAlbumById(Album{
		ID: 1,
		Title: "Red Brain",
		Artist: "John Coltrane",
		Price: 99.99,
	}, db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of updated album: %v\n", newAlbID)

	row, err := deleteAlbumById(5, db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of deleted album: %v\n", row)
}

// db is not initialized if no db parameter put in
func fetchAlbumsByArtist(name string, db *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)

	if err != nil {
		return nil, fmt.Errorf("fetchAlbumByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

	return albums, nil
}

func addAlbum(alb Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("insert into album (title, artist, price) values (?, ?, ?)", alb.Title, alb.Artist, alb.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func updateAlbumById(newAlb Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("update album set title = ?, artist = ?, price = ? where id = ?", newAlb.Title, newAlb.Artist, newAlb.Price, newAlb.ID)

	if err != nil {
		return 0, fmt.Errorf("updateAlbumById: %v", err)
	}

	id, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("updateAlbumById: %v", err)
	}

	return id, nil
}

func deleteAlbumById(id int, db *sql.DB) (int64, error) {
	result, err := db.Exec("delete from album where id = ?", id)

	if err != nil {
		return 0, fmt.Errorf("deleteAlbumById: %v", err)
	}

	row, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("deleteAlbumById: %v", err)
	}

	return row, nil
}
```

## 6. How to use viper to read config file

```go
type Env struct {
	AccessTokenExpiryHour int `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`		// no space allowed
	RefreshTokenExpiryHour int `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

func main() {
	env := NewEnv()

	fmt.Println(env.AccessTokenExpiryHour)
	fmt.Println(env.AccessTokenSecret)
}
```

## 7. How to use swaggo for generating api docs

```go
package route

import (
	"backend/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewSwagDocRouter(env *bootstrap.Env, group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

ATTENTION: swagger's auto import is not very good, so developers need to import these dependencies
manually.

Also, in main.go, or entry file, add this import:

```go
package main

import (
    _ "yourmodule/docs"
)
```

This is because swaggo will use files in folder docs in order to generate api docs.