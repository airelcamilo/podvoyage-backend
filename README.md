# Podvoyage Backend

Welcome to the backend code repository for Podvoyage, a podcast player application. With Podvoyage, you can easily search your favorite podcasts, save it for later, and listen to it. This project utilizes `Golang`, accompanied with packages, such as `gorm`, `mux`, and `cors`. Also this project uses `Supabase` for the database system and deployed using `Google Cloud Run`.

## Features
1. **Create or login** to your account
2. **Search your favorite podcast** using iTunes API
3. **Save or remove podcast** from the databse
4. **Create or remove folder** for organizing your podcasts
5. **Mark podcasts as played**
6. **Resuming podcast** from where you left off
7. **Add or remove podcast from queue** (Not integrated with frontend yet)

## Getting Started

Using local `postgreSQL` database, in `db.go` change 
```
os.Getenv("SUPABASE_DB") -> "host=localhost user=postgres password=postgres dbname=podvoyage port=5432"
```

Then, add the `podvoyage` database:
```SQL
CREATE DATABASE podvoyage;
```

With your database in place, initiate the development server:

```bash
cd .\cmd\podvoyage\
go run main.go
```

Open [http://localhost:4000](http://localhost:4000) and you can start interacting with the API.

In this repository, there also provided a comprehensive Postman collection:
1. [Postman Podvoyage JSON collection](podvoyage_postman.json)
2. [Postman User JSON collection](user_postman.json)

## Available API Calls

<details>
<summary>User</summary>
<br>

* `[POST]` Register
    
    Url: 
    ```
    /api/register
    ```
    Body:
    ```JSON
    {
      "name": string,
      "email": string,
      "username": string,
      "password": string
    }
    ```
    Return: 
    ```JSON
    {
      "token": string,
      "user": User
    }
    ```

* `[POST]` Login

    Url: 
    ```
    /api/login
    ```
    Body:
    ```JSON
    {
      "email": string,
      "password": string
    }
    ```
    Return: 
    ```JSON
    {
      "token": string,
      "user": User
    }
    ```

* `[POST]` Validate

    Url: 
    ```
    /api/validate
    ```
    Body:
    ```JSON
    {
      "token": string,
    }
    ```
    Return: 
    ```JSON
    {
      "name": string,
      "email": string,
      "username": string,
      "password": string
    }
    ```

* `[POST]` Logout
    
    Url: 
    ```
    /api/logout
    ```
    Body:
    ```JSON
    {
      "token": string,
    }
    ```
    Return: 
    ```JSON
    string
    ```
</details>

<details>
<summary>Podcast</summary>
<br>

* `[POST]` Search All Podcast
    
    Url: 
    ```
    /api/search-all
    ```
    Body:
    ```JSON
    {
      "podcastName": string
    }
    ```
    Return: 
    ```JSON
    {
      "resultCount": int,
      "results": []Podcast
    }
    ```

* `[GET]` Search Podcast

    Url: 
    ```
    /api/search-pod/{trackId}
    ```
    Return:
    ```JSON
    {
      "id": int,
      "trackId": int,
      "trackName": string,
      "artistName": string,
      "feedUrl": string,
      "artworkUrl600": string,
      "desc": string,
      "link": string,
      "categories": []Category,
      "episodes": []Episode
    }
    ```

* `[GET]` Get All Podcast

    Url: 
    ```
    /api/podcasts
    ```
    Return: 
    ```JSON
    []Podcast
    ```

* `[GET]` Get Podcast
    
    Url: 
    ```
    /api/podcast/{podId}
    ```
    Return: 
    ```JSON
    {
      "id": int,
      "trackId": int,
      "trackName": string,
      "artistName": string,
      "feedUrl": string,
      "artworkUrl600": string,
      "desc": string,
      "link": string,
      "categories": []Category,
      "episodes": []Episode
    }
    ```

* `[POST]` Save Podcast
    
    Url: 
    ```
    /api/podcast
    ```
    Body:
    ```JSON
    {
      "id": int,
      "trackId": int,
      "trackName": string,
      "artistName": string,
      "feedUrl": string,
      "artworkUrl600": string,
      "desc": string,
      "link": string,
      "categories": []Category,
      "episodes": []Episode
    }
    ```
    Return: 
    ```JSON
    {
      "id": int,
      "trackId": int,
      "trackName": string,
      "artistName": string,
      "feedUrl": string,
      "artworkUrl600": string,
      "desc": string,
      "link": string,
      "categories": []Category,
      "episodes": []Episode
    }
    ```
    
* `[DELETE]` Remove Podcast
    
    Url: 
    ```
    /api/podcast/{podId}
    ```
    Return: 
    ```JSON
    podId: int
    ```
</details>

<details>
<summary>Item</summary>
<br>

* `[GET]` All Item
    
    Url: 
    ```
    /api/all
    ```
    Return: 
    ```JSON
    [
      {
        "id": int,
        "type": string,
        "name": string,
        "artistName": string,
        "artworkUrl": string,
        "podcastId": int,
        "trackId": int,
        "folderId": int,
        "pos": int
      }
    ]
    ```
</details>

<details>
<summary>Folder</summary>
<br>

* `[GET]` Get All Folder
    
    Url: 
    ```
    /api/folders
    ```
    Return: 
    ```JSON
    []Folder
    ```


* `[GET]` Get Folder

    Url: 
    ```
    /api/folder/{folderId}
    ```
    Return:
    ```JSON
    {
      "id": int,
      "folderName": string,
      "podcasts": []Podcast
    }
    ```

* `[POST]` Save Folder
    
    Url: 
    ```
    /api/folder
    ```
    Body:
    ```JSON
    {
      "folderName": string
    }
    ```
    Return: 
    ```JSON
    {
      "id": int,
      "folderName": string,
      "podcasts": []Podcast
    }
    ```

* `[GET]` Check in Folder
    
    Url: 
    ```
    /api/in-folder/{podId}
    ```
    Return: 
    ```JSON
    folderId: int
    ```

* `[GET]` Change Folder
    
    Url: 
    ```
    /api/change-folder/{folderId}/{podId}
    ```
    Return: 
    ```JSON
    folderId: int
    ```
    
* `[DELETE]` Remove Folder
    
    Url: 
    ```
    /api/folder/{folderId}
    ```
    Return: 
    ```JSON
    folderId: int
    ```
</details>

<details>
<summary>Queue</summary>
<br>

* `[GET]` Get All Queue
    
    Url: 
    ```
    /api/queue
    ```
    Return: 
    ```JSON
    [
      {
      "episode": Episode,
      "episodeId": int,
      "pos": int
      }
    ]
    ```


* `[POST]` Add to Queue

    Url: 
    ```
    /api/queue
    ```
    Body:
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```
    Return:
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```

* `[DELETE]` Remove in Queue
    
    Url: 
    ```
    /api/queue/{episodeId}
    ```
    Return: 
    ```JSON
    episodeId: int
    ```
</details>

<details>
<summary>Episode</summary>
<br>

* `[POST]` Mark as Played
    
    Url: 
    ```
    /api/played/{episodeId}
    ```
    Body:
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```
    Return: 
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```


* `[POST]` Set Current Time

    Url: 
    ```
    /api/folder/{id}
    ```
    Body:
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```
    Return: 
    ```JSON
    {
      "id": int,
      "podcastId": int,
      "trackId": int,
      "title": string,
      "desc": string,
      "season": int,
      "date": string,
      "duration": int,
      "audio": string,
      "played": bool,
      "currentTime": int
    }
    ```
</details>


## Credit
Airel Camilo Khairan © 2023