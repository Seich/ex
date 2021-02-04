# Coding Test

## Building and Running
```
git clone git@github.com:Seich/ex.git
cd ex
go run .
```

## Usage

The `ex` service exposes 4 endpoints on your local `12933` port:

### `GET /students`
This endpoint returns a list of all student ids with existing exam results.

```sh
$ curl http://0.0.0.0:12933/students
```

### `GET /student/:studentId`
This endpoint receives one of the student ids from the previous request. It will
return the Student, their exam results and their average score.

```sh
# Amari26 is one of the StudentIds that came as part of the previous request.
$ curl http://0.0.0.0:12933/student/Amari26 
```

### `GET /exams`
This endpoint returns a list of all exam ids that have been taken by the students.

```sh
# Amari26 comes from the list on the previous request.
$ curl http://0.0.0.0:12933/exams
```

### `GET /exam/:examId`
This endpoint returns all existing exam results, along with the average test
score for the given exam.

```sh
# 1234 comes from the list on the previous request.
$ curl http://0.0.0.0:12933/exam/1234
```

A [`beau.yml`](https://beaujs.com/) file is included in the repo, if beau is installed (`npm install -g
beau`) you can just call `beau request -i` to walk through the different
requests.

## Dependencies
The only dependency being used is http://github.com/tidwall/buntdb which is used
both to demonstrate the use of an external library and as the in-memory data
store for the server. This dependency should be automatically installed by go on
first build.
