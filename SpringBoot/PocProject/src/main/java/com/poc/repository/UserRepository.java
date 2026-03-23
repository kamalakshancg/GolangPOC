package com.poc.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import com.poc.entity.UserEntity;

import java.util.Set;

public interface UserRepository extends JpaRepository<UserEntity, Integer> {
    @Query("SELECT DISTINCT u FROM UserEntity u " +
           "LEFT JOIN FETCH u.orders o " +
           "LEFT JOIN FETCH o.items " +
           "WHERE u.id BETWEEN 1 AND 50")
    Set<UserEntity> findUsersWithOrdersAndItems();
}