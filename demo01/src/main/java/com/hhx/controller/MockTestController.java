package com.hhx.controller;

import com.hhx.service.IMockTestService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

// UserController.java
@RestController
@RequestMapping("/mocktest")
public class MockTestController {


    @Autowired
    private IMockTestService mockTestService;

    @RequestMapping("/get")
    public String get() {
        return mockTestService.get();
    }
}
