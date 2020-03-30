package com.zx5435.wolan;

import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@RestController
@RequestMapping("/other")
@CrossOrigin
public class OtherController {

    @RequestMapping("aaa")
    public Object qwe() {
        HashMap<Object, Object> m = new HashMap<>();
        m.put("name", "asd");
        return m;
    }

}
