package com.poc.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.poc.entity.OrderEntity;

import org.springframework.data.domain.Pageable;
import java.util.List;

public interface OrderRepository extends JpaRepository<OrderEntity, Integer> {
    List<OrderEntity> findBy(Pageable pageable);
}