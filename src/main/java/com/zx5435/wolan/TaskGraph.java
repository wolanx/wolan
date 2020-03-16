package com.zx5435.wolan;

import com.coxautodev.graphql.tools.GraphQLQueryResolver;
import com.zx5435.wolan.model.TaskDO;
import org.springframework.stereotype.Component;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;


@Component
public class TaskGraph implements GraphQLQueryResolver {

    public TaskDO getTaskByName(String name) {
        return null;
    }

    public List<TaskDO> listTask() {
        ArrayList<TaskDO> res = new ArrayList<>();
        System.out.println("res = " + res);

        File file = new File("./src/main/resources/gitops");
        File[] files = file.listFiles();

        for (File task : Objects.requireNonNull(files)) {
            if (task.isDirectory()) {
                TaskDO taskDO = new TaskDO();
//                taskDO
                res.add(taskDO);
            }
        }

        return res;
    }

}
