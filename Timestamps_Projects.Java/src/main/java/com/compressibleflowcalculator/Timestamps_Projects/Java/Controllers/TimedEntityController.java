package com.compressibleflowcalculator.Timestamps_Projects.Java.Controllers;


import java.time.OffsetDateTime;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

//REST API Controller at /api/model_with_Timestamp

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.compressibleflowcalculator.Timestamps_Projects.Java.Models.ModelWithTimestamp;
import com.compressibleflowcalculator.Timestamps_Projects.Java.Models.Notification;
import com.compressibleflowcalculator.Timestamps_Projects.Java.Repository.NotificationRepository;
import com.compressibleflowcalculator.Timestamps_Projects.Java.Repository.PostgresRepository;

import org.springframework.web.bind.annotation.RequestBody;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;

import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;

@RestController
@RequestMapping("/api/model_with_Timestamp")
public class TimedEntityController {
    
    //The Model
    @Autowired
    private PostgresRepository postgresRepository;

    @Autowired
    private NotificationRepository notificationRepository;

    //constructor
    public TimedEntityController(PostgresRepository postgresRepository, NotificationRepository notificationRepository) {
        this.postgresRepository = postgresRepository;
        this.notificationRepository = notificationRepository;
    }

    //Post Request
    @PostMapping
    @Operation(summary = "Create a new TimedEntity")
    @ApiResponses(value = {
        @ApiResponse(responseCode = "201", description = "Created the TimedEntity",
            content = { @Content(mediaType = "application/json",
            schema = @Schema(implementation = ModelWithTimestamp.class)) }),
        @ApiResponse(responseCode = "400", description = "Invalid input",
            content = @Content)
    })
    public ResponseEntity<ModelWithTimestamp> createTimedEntity(@RequestBody ModelWithTimestamp timedEntity) {

        System.out.println(timedEntity);
        System.out.println(timedEntity.getName());
        System.out.println(timedEntity.getTimestamp());
        if(timedEntity.getName() == null || timedEntity.getName().isEmpty()) {
           return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
        }
        ModelWithTimestamp savedEntity = postgresRepository.save(timedEntity);

        //CREATE OffsetTImeDate That is 3 hours Behind timedEntity.getTimestamp()
        OffsetDateTime offsetDateTime = timedEntity.getTimestamp().minusHours(3);

        //create new notification with this offsetDateTime
        notificationRepository.save(new Notification(timedEntity.getName(), offsetDateTime));

        return new ResponseEntity<>(savedEntity, HttpStatus.CREATED);

    }
    

    //Get Request
    @GetMapping
    public ResponseEntity<Iterable<ModelWithTimestamp>> getTimedEntities() {
        Iterable<ModelWithTimestamp> timedEntities = postgresRepository.findAll();

        // Iterate THrough the Entities
        for(ModelWithTimestamp entity : timedEntities) {
            System.out.println(entity.getName());
            System.out.println(entity.getTimestamp());
        }

        return new ResponseEntity<>(timedEntities, HttpStatus.OK);
    }


}
