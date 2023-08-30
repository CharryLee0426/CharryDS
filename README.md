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

## 8. Concurrency

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	go function()

	time.Sleep(1 * time.Second)
}

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
```

Execute results I tested:
```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
10 11 12 13 14 15 16 17 18 19 0123456789
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
10 11 12 13 14 15 16 17 18 19 0123456789
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
10 11 12 13 14 15 16 17 18 19 0123456789
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
0123456789
10 11 12 13 14 15 16 17 18 19 % 
```

It shows that developers can't control the order of execution of goroutines.

Another example:

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
```

```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go -n 100
Going to create 100 goroutines.
0 50 3 4 2 27 5 28 30 31 6 29 7 8 9 10 11 12 32 13 33 14 15 17 16 99 51 34 35 37 20 36 53 40 55 23 41 54 39 26 57 22 19 52 48 24 56 61 18 42 71 21 47 73 45 38 65 59 60 76 25 80 69 49 63 87 81 82 72 83 89 91 84 58 92 75 85 46 66 67 68 62 77 78 86 79 43 93 70 88 96 44 97 98 64 94 74 90 95 1 
Exiting...
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go -n 100
Going to create 100 goroutines.
3 2 23 14 15 16 17 18 19 20 21 22 8 1 0 10 4 5 6 7 9 34 24 46 25 35 26 28 29 30 37 36 27 38 39 40 41 31 42 43 32 44 33 45 72 11 12 85 75 74 73 47 48 49 52 51 50 99 80 77 78 79 88 56 83 54 89 55 76 91 82 86 60 87 84 97 57 90 62 81 95 64 53 92 96 93 98 94 67 66 68 61 58 69 71 70 59 65 63 13 
Exiting...
```

Another point needs to pay attention to: each program needs a sleep time, otherwise, the program will exit before goroutines are executed.

Another way for goroutines to execute successfully is to use `sync.WaitGroup`:

```go
package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("\n%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
```

We have to put Add operation first in order to prevent race condition.

Result:
```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go -n 100
Going to create 100 goroutines.
sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0x0}, sema:0x0}
0 8 5 6 7 54 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 2 1 3 23 9 10 11 12 13 14 15 16 17 18 19 20 21 22 70 55 56 57 4 63 58 64 
sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0x3900000000}, sema:0x0}
59 65 61 60 99 69 68 67 62 66 71 76 85 73 25 75 79 72 90 24 28 82 77 96 78 98 86 34 30 27 80 81 87 91 33 74 29 93 88 97 83 89 94 95 84 26 36 32 35 31 37 92 38 
Exiting...
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go -n 100
Going to create 100 goroutines.
sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0x0}, sema:0x0}
3 0 14 4 5 6 7 8 9 10 11 12 13 26 15 16 17 18 19 20 21 22 23 24 25 40 27 1 33 28 34 35 30 29 36 37 31 38 32 65 41 2 
sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0x3b00000000}, sema:0x0}
82 43 66 47 67 42 85 45 50 74 81 51 46 44 71 88 91 84 72 78 48 57 97 92 69 99 39 94 75 68 86 76 95 87 83 52 96 53 54 55 77 56 90 89 49 61 79 58 73 62 59 80 63 60 64 93 70 98 
Exiting...
```

If `waitGroup.Add` > `waitGroup.Done`, the program will return an error "fatal error: all goroutines are asleep - deadlock!";
If `waitGroup.Add` < `waitGroup.Done`, the program will return an error "panic: sync: negative WaitGroup counter".

Channel is a mechanism for goroutines to communicate with each other. It is a data structure that can be used to send and receive messages between goroutines. In practice, channels should be specified to be used
for just one purpose, either sending or receiving if they're used as
parameters.

Read this example:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x

    // ignored because no one reads contents in channel
	close(c)
	fmt.Println(x)
}
```

Golang uses arrows to show the direction of data flow. It's very easy
for developers to understand.

```go
// c is a chan, v is a value, k is a variable
// c, v, and k have the same value type

// write
c <- v

// read
k := <- c
```

Fix the above example to let ignored part be executed:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <- c

	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}
```

Result:
```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
1 10
Read: 10
2 10
Channel is closed!
```

Pipeline is a system that creates a constant data flow in programms. It's a series of stages connected by channels, where each stage is a group of goroutines running the same function.

```go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false				// checked by first(), altered bt second()
var DATA = make(map[int]bool)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())

	A := make(chan int) // for input
	B := make(chan int) // for output

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}

func random(min, max int) int {
	return min + rand.Intn(max - min)
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")

		_, ok := DATA[x]

		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}

	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	sum := 0

	for x := range in {
		sum += x
	}

	fmt.Printf("The sum of the random numbers is %d.\n", sum)
}
```

Practice: implement a pipeline that generates sum of squares of numbers in a range by both sequential and concurrent ways.

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be less than %d\n", n1, n2)
	}

	// sum := 0

	// for i := n1; i <= n2; i++ {
	// 	sum += i*i
	// }

	// fmt.Println(sum)

	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(A, B)
	third(B)

	fmt.Println(time.Since(start))
}

func first(min, max int, out chan<- int) {
	i := min

	for i <= max {
		out <- i
		i++
	}

	close(out)
}

func second(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}

	close(out)
}

func third(in <-chan int) {
	sum := 0

	for x := range in {
		sum += x
	}

	fmt.Println(sum)
}
```

Result: sequential implementation is faster than concurrent. (50us vs 150us) It makes sense because concurrent implementation contains extra operations such as creating and close channels, creating goroutines, etc.

Using `select` keyword can connect several channels. It is very important in go concurrency model. The usage of `select` is very similar
to `switch` statement. However, it checks simultaneously but not sequentially.

```go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)
	go gen(0, 2*n, createNumber, end)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true
}

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max - min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}
```

It's very important to learn how to **time out** goroutines. There are two ways to do it. Let me show the first way:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		fmt.Println("c1 start")
		time.Sleep(time.Second * 3) // simulate this task needs 3 seconds
		c1 <- "c1 is ok"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)

	go func() {
		fmt.Println("c2 start")
		time.Sleep(time.Second * 3) // simulate this task needs 3 seconds
		c2 <- "c2 is ok"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 4):
		fmt.Println("timeout c2")
	}
}
```

You will see c1 is time out but c2 is executed successfully.

```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
timeout c1
1.001023166s
c2 is ok
4.00227725s
4.002290333s
```

Another way is to use `sync.WaitGroup` to control the time out:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	arguments := os.Args
	start := time.Now()

	if len(arguments) != 2 {
		fmt.Println("Need a time duration!")
		return
	}

	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is: %d\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
		fmt.Println(time.Since(start))
	} else {
		fmt.Println("OK!")
		fmt.Println(time.Since(start))
	}

	w.Done()

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
		fmt.Println(time.Since(start))
	} else {
		fmt.Println("OK!")
		fmt.Println(time.Since(start))
	}
}

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)

	go func() {
		time.Sleep(5 *time.Second)
		defer close(temp)
		w.Wait()
	}()

	select {
		case  <-temp:
			return false
		case <-time.After(t):
			return true
	}
}
```

If we set duration to 10s, the result should be:
```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go 10000
Timeout period is: 10000000000
Timed out!
10.001309208s
OK!
15.002601167s
```

If we set duration to 1s, the result should be:
```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go 1000
Timeout period is: 1000000000
Timed out!
1.001271042s
Timed out!
2.002877084s
```

A buffered channel can be used as semaphore in order to limit the throughput of your application.

```go
package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("No more numbers")
			break
		}
	}
}
```
In this program, only 0 to 4 can be put into the buffered channel.

`nil` channel is useful due to its always blocked feature.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func  main() {
	c := make(chan int)

	go add(c)
	go send(c)

	time.Sleep(2 * time.Second)
}

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil		// it will be blocked
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}
```

Channel of channels is very hard for developers to understand. This is an example to illustrate how to use it. Hope me can understand this:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var times int

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("Need just one argument!")
		return
	}

	times, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	cc := make(chan chan int)

	for i := 1; i < times+1; i++ {
		f := make(chan bool)

		go f1(cc, f)

		ch := <-cc
		ch <- i

		for sum := range ch {
			fmt.Printf("Sum of %d is %d\n", i, sum)
		}

		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)

	cc <- c
	defer close(c)

	sum := 0

	select {
	case x := <-c:
		for i := 0; i < x; i++ {
			sum += i
		}

		c <- sum
	case <-f:
		return
	}
}
```

I can know its usage. The real problem is: how this program can make sure the order of `c := make(chan int)`, `cc <- c`, `ch := <-cc`, and `ch <- i`?

Critical Section is very important when you want some concurrency tasks runs in order. For example, you should do this if many tasks share the same memory or variables. To achive it, you can use `sync.Mutex`.

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m sync.Mutex
	v1 int
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%d", read())

	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)

		go func(x int) {
			defer waitGroup.Done()
			change(x)
			fmt.Printf("-> %d", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
}

func change(i int) {
	m.Lock()

	// this section will be executed by only one goroutine at a time
	// this section is called critical section
	time.Sleep(time.Second)
	v1 += 1
	if v1 % 10 == 0 {
		v1 -= 10 * i
	}

	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}
```

If you forget to unlock your critical section, the program will be deadlocked.

For eample, this program:

```go
package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func main() {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()

	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()

	w.Add(1)

	w.Wait()
}

func function() {
	m.Lock()
	fmt.Println("Locked")
	// no unlock
}
```

```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go
Locked
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x140000021a0?)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/sema.go:62 +0x2c
sync.(*WaitGroup).Wait(0x140000a6020)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/sync/waitgroup.go:116 +0x78
main.main()
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:27 +0xd0

goroutine 18 [sync.Mutex.Lock]:
sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/sema.go:77 +0x28
sync.(*Mutex).lockSlow(0x10427ff40)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/sync/mutex.go:171 +0x178
sync.(*Mutex).Lock(...)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/sync/mutex.go:90
main.function()
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:31 +0x84
main.main.func1()
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:15 +0x4c
created by main.main
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:13 +0x6c
exit status 2
```

`RWMutex` is more advanced exclusive mutex. It allows multiple readers or one writer at the same time. Its drawback is that it is only when there are no readers reading the data, the writer can write data.

```go
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM sync.RWMutex
	M sync.Mutex
	password string
}

var password = secret{password: "myPassword"}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)

	c.password = pass
	c.RWM.Unlock()
}

// Using RWMutex
func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show")

	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	
	return c.password
}

// Using Mutex
func showWithLock(c *secret) string {
	c.M.Lock()

	fmt.Print("showWithLock")

	time.Sleep(3 * time.Second)
	defer c.M.Unlock()

	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }

	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}

	var waitGroup sync.WaitGroup

	fmt.Println("Pass:", showFunction(&password))

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)

		// reader
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", showFunction(&password))
		}()

		// writer
		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			Change(&password, "123456")
		}()

		waitGroup.Wait()
		fmt.Println("Pass:", showFunction(&password))
	}
}
```

Result here:
```
(base) chenli@Chens-MacBook-Pro CharryDS % time go run main.go 10  >/dev/null
go run main.go 10 > /dev/null  0.13s user 0.14s system 0% cpu 1:08.76 total
(base) chenli@Chens-MacBook-Pro CharryDS % time go run main.go  >/dev/null 
go run main.go > /dev/null  0.12s user 0.14s system 0% cpu 1:23.43 total
```

The first one used `Mutex` and the second one used `RWMutex`. The result shows that `RWMutex` is faster than `Mutex`.

A data race condition is a situation where two or more running elements such as threads and goroutines try to take control or modify a shared resource or a variable of a program. A data race condition occurs when there are two or more instuctions access the same memory and at least one of them is a write instruction.

Using `-race` flag when running or building a Go source file.

For example, read this program:
```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide an integer")
		os.Exit(1)
	}

	numGR, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	k := make(map[int]int)
	k[1] = 12

	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			k[i] = i
		}(i)
	}

	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}
```

```
==================
WARNING: DATA RACE
Write at 0x00c000100180 by goroutine 9:
  runtime.mapaccess2_fast64()
      /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/map_fast64.go:53 +0x1cc
  main.main.func1()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:32 +0x78
  main.main.func2()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:33 +0x44

Previous write at 0x00c000100180 by goroutine 6:
  runtime.mapaccess2_fast64()
      /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/map_fast64.go:53 +0x1cc
  main.main.func1()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:32 +0x78
  main.main.func2()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:33 +0x44

Goroutine 9 (running) created at:
  main.main()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:30 +0x1cc

Goroutine 6 (finished) created at:
  main.main()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:30 +0x1cc
==================
==================
WARNING: DATA RACE
Write at 0x00c0000260f8 by main goroutine:
  main.main()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:36 +0x30c

Previous write at 0x00c0000260f8 by goroutine 8:
  main.main.func1()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:32 +0x84
  main.main.func2()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:33 +0x44

Goroutine 8 (finished) created at:
  main.main()
      /Users/chenli/Documents/ComputerScience/CharryDS/main.go:30 +0x1cc
==================
fatal error: concurrent map writes

goroutine 11 [running]:
main.main.func1(0x5)
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:32 +0x7c
created by main.main
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:30 +0x1d0

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000014108?)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/sema.go:62 +0x2c
sync.(*WaitGroup).Wait(0xc000014100)
        /opt/homebrew/Cellar/go/1.20.6/libexec/src/sync/waitgroup.go:116 +0x7c
main.main()
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:37 +0x324

goroutine 14 [runnable]:
main.main.func1(0x8)
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:32 +0x7c
created by main.main
        /Users/chenli/Documents/ComputerScience/CharryDS/main.go:30 +0x1d0
exit status 2
```

How to solve race condition? It's easy, just adding mutex.

`Context` package is simple but useful. It is used to do some cancellation or timeout operations.

Here is an example program:

```go
package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}

	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t) * time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}

	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2 * t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}

	return
}
```

The result:

```
(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go 4
Delay: 4
f1(): 2023-08-30 00:12:07.227739 -0700 PDT m=+4.001223960
f2(): 2023-08-30 00:12:11.229846 -0700 PDT m=+8.003349751
f3(): 2023-08-30 00:12:15.231034 -0700 PDT m=+12.004555376

(base) chenli@Chens-MacBook-Pro CharryDS % go run main.go 10
Delay: 10
f1(): context canceled
f2(): context canceled
f3(): context canceled
```