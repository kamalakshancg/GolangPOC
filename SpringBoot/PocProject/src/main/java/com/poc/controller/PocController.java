package com.poc.controller;

import com.poc.dto.UserWithOrders;
import com.poc.entity.Order;
import com.poc.entity.User;
import com.poc.services.OrderService;
import com.poc.services.UserService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;
import java.util.Set;

import com.poc.repository.OrderRepository;
import com.poc.repository.UserRepository;

@RestController
public class PocController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    @Autowired
    private OrderService orderService;

    @Autowired
    private UserService userService;

    @GetMapping("/api/test1")
    public String test1() {
        return jdbcTemplate.queryForObject("SELECT 'pong'", String.class);
    }

    @GetMapping("/api/test2")
    public List<Order> test2() {
        return orderService.getOrderDetails();
    }

    @GetMapping("/api/test3")
    public List<User> test3() {
        return userService.userWithOrder();
    }
}