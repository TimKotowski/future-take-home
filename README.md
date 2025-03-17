# future-take-home-test

## Requirements To Run
1. Docker installed and able to interact with docker.

## Steps to Run Take home test

```cgo
docker compose up
```
## Fully passing take home with listed requirements

- [x] The client should be able to pick from a list of available times, and appointments
  for a coach should not overlap.
- [x] All appointments are 30 minutes long, and should be scheduled at :00, :30
  minutes after the hour during business hours. 
- [x] Business hours are M-F 8am-5pm Pacific Time
- [x] All API's covered
  - Get a list of available appointment times for a trainer between two dates
  - Post a new appointment
  - Get a list of scheduled appointments for a trainer
- [x] Using Postgres
- [x] Dockerfile or equivalent for the reviewer to follow setup steps
- [x] Test Coverage
  - Test Coverage is not fully done, due to the time constraints of the take home. It was recommended not to spend more than 3 hours on this (went little over and spent 4 hours on the take home assignment)

## Testing API
`All Test data is seeded in migrations script based on JSON data provided for test`

- Get a list of available appointment times for a trainer between two dates
  - ``http://localhost:8080/appointments/v1/slots/:trainerID``
    - request body 
    ```
    start_slot
    end_slot
    ```

- Get a list of scheduled appointments for a trainer
  - ``http://localhost:8080/appointments/v1/slots/:trainerID``

- Post a new appointment
  - ``http://localhost:8080/appointments/v1/slots/:trainerID/:userID``
    - request body
    ```
    start_slot
    end_slot
    status
    ```