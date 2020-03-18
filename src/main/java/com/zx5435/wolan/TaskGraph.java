package com.zx5435.wolan;

import com.coxautodev.graphql.tools.GraphQLQueryResolver;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.util.JSONPObject;
import com.zx5435.wolan.model.TaskDO;
import org.springframework.boot.configurationprocessor.json.JSONObject;
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

    private static final String WorkPath = "./src/main/resources/gitops";

    public TaskDO getTaskByName(String taskName) throws FileNotFoundException {
        Yaml yaml = new Yaml();
        File woFile = new File(WorkPath + "/" + taskName + "/wolan.yaml");
        Object obj = yaml.load(new FileInputStream(woFile));

        ObjectMapper mapper = new ObjectMapper();
        TaskDO task = mapper.convertValue(obj, TaskDO.class);

        System.out.println("task = " + task);
        return task;
    }

    public List<TaskDO> listTask() throws IOException {
        ArrayList<TaskDO> res = new ArrayList<>();

        File workFile = new File(WorkPath);
        File[] taskFiles = workFile.listFiles();

        for (File taskFile : Objects.requireNonNull(taskFiles)) {
            if (taskFile.isDirectory()) {
                String taskName = taskFile.getName();
                TaskDO task = getTaskByName(taskName);
                res.add(task);
            }
        }

        return res;
    }

}
