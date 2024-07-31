package com.compressibleflowcalculator.Timestamps_Projects.Java;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.web.servlet.support.AbstractAnnotationConfigDispatcherServletInitializer;


@Configuration
@EnableWebSecurity
public class SecurityChainFilter {

    @Bean
    public SecurityFilterChain web(HttpSecurity http) throws Exception {
        http
            .csrf().disable()
            .authorizeHttpRequests((authorize) -> authorize
            .requestMatchers("/*", "/*/**").permitAll());   
        return http.build();
    }


}