package com.zx5435.wolan.graph;

import com.coxautodev.graphql.tools.GraphQLQueryResolver;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;
import org.springframework.stereotype.Component;
import org.yaml.snakeyaml.Yaml;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;


@Component
public class TaskGraph implements GraphQLQueryResolver {

    public static TaskDO getTaskByName(String sid) throws IOException {
        Yaml yaml = new Yaml();
        File f = new File(WoConf.WorkPath + "/" + sid + "/wolan.yml");
        FileInputStream fIn = new FileInputStream(f);
        Object obj = yaml.load(fIn);
        fIn.close();

        ObjectMapper mapper = new ObjectMapper();
        TaskDO task = mapper.convertValue(obj, TaskDO.class);
        task.setSid(sid);

        System.out.println("task = " + task);
        return task;
    }

    public List<TaskDO> listTask() throws IOException {
        ArrayList<TaskDO> res = new ArrayList<>();

        File workFile = new File(WoConf.WorkPath);
        File[] taskFiles = workFile.listFiles();

        for (File taskFile : Objects.requireNonNull(taskFiles)) {
            if (taskFile.isDirectory()) {
                res.add(getTaskByName(taskFile.getName()));
            }
        }

        return res;
    }

}
