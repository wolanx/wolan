package com.zx5435.wolan.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import java.util.HashMap;

@Controller
public class OtherController {

    @RequestMapping("/")
    public String index() {
        System.out.println("true = " + true);
        return "site/index";
    }

    @RequestMapping("/aaa")
    public Object test() {
        HashMap<Object, Object> m = new HashMap<>();
        m.put("data", new int[]{22, 3, 4, 5, 7, 1});
        return m;
    }

}
