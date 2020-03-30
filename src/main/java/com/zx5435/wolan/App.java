package com.zx5435.wolan;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.web.bind.annotation.CrossOrigin;

@SpringBootApplication
public class App {

    public static void main(String[] args) {
        new SpringApplicationBuilder(App.class).run(args);
    }

}
