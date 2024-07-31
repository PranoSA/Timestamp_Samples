package com.compressibleflowcalculator.Timestamps_Projects.Java.Repository;

import java.util.UUID;

import org.springframework.data.repository.CrudRepository;

import com.compressibleflowcalculator.Timestamps_Projects.Java.Models.Notification;

public interface NotificationRepository extends CrudRepository<Notification, UUID> {
    
    
}
