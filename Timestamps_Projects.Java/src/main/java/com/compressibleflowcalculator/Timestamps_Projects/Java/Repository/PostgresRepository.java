package com.compressibleflowcalculator.Timestamps_Projects.Java.Repository;

import org.springframework.data.repository.CrudRepository;

import com.compressibleflowcalculator.Timestamps_Projects.Java.Models.ModelWithTimestamp;

public interface PostgresRepository extends CrudRepository<ModelWithTimestamp, String> {
}
