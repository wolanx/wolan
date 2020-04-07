package com.zx5435.wolan;

import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@RestController
@RequestMapping("/other")
@CrossOrigin
public class OtherController {

    @RequestMapping("test")
    public Object test() {
        HashMap<Object, Object> m = new HashMap<>();
        m.put("data", new int[]{22, 3, 4, 5, 7, 1});
        return m;
    }

}
