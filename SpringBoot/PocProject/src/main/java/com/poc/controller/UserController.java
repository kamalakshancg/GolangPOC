package com.poc.controller;

import com.poc.entity.User;
import com.poc.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserService userService;

    @GetMapping("/getUserWithOrders")
    public List<User> getUserWithOrders() {
        return userService.getUserWithOrderDetails();
    }
}
