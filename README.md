# temporal-helloworld

## Bring the Temporal cluster up
  ```
  git clone https://github.com/temporalio/docker-compose.git
  cd docker-compose
  docker-compose up
  ```

## Run the worker
In a separate terminal window.
```
go run worker/main.go
```
![Screen Shot 2022-09-05 at 4 29 19 PM](https://user-images.githubusercontent.com/43081882/188513985-a594b92b-9c9c-446b-bd62-caaa6da3b3e7.png)



## Run the starter
In a separate terminal window.
```
go run starter/main.go
```

![Screen Shot 2022-09-05 at 4 29 07 PM](https://user-images.githubusercontent.com/43081882/188514009-d67ec3f5-0d4f-4b62-9d18-e42c235d98f1.png)
