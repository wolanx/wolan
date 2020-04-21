package com.zx5435.wolan.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;
import lombok.SneakyThrows;
import org.apache.commons.io.FileUtils;
import org.yaml.snakeyaml.Yaml;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileWriter;
import java.io.IOException;

public class TaskService {

    @SneakyThrows
    public static TaskDO getOne(String sid) {
        Yaml yaml = new Yaml();
        File f = new File(WoConf.WorkPath + "/" + sid + "/wolan.yml");
        FileInputStream fIn = new FileInputStream(f);
        Object obj = yaml.load(fIn);
        fIn.close();

        ObjectMapper mapper = new ObjectMapper();
        TaskDO task = mapper.convertValue(obj, TaskDO.class);
        task.setSid(sid);

        return task;
    }

    @SneakyThrows
    public static TaskDO addOne(TaskDO task) {
        task.setSid(task.getName());
        task.setVersion("1.0.0");

        File dir = new File(WoConf.WorkPath + "/" + task.getSid());
        if (!dir.exists()) {
            boolean b = dir.mkdir();
        }

        File f = new File(WoConf.WorkPath + "/" + task.getSid() + "/wolan.yml");
        if (!f.exists()) {
            boolean b = f.createNewFile();
        }
        System.out.println("task = " + task);
        System.out.println("dir.exists() = " + dir.exists());
        System.out.println("f.exists() = " + f.exists());

        FileWriter fileWriter = new FileWriter(f);
        Yaml yaml = new Yaml();
        yaml.dump(task, fileWriter);
        fileWriter.close();
        return task;
    }

    public static TaskDO updateOne(TaskDO task) {
        return task;
    }

    public static Boolean deleteOne(String sid) {
        File dir = new File(WoConf.WorkPath + "/" + sid);

        try {
            FileUtils.deleteDirectory(dir);
            return true;
        } catch (IOException e) {
            return false;
        }
    }

}
