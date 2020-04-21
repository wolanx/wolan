package com.zx5435.wolan.graph;

import com.coxautodev.graphql.tools.GraphQLMutationResolver;
import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.service.TaskService;
import org.springframework.stereotype.Component;


@Component
public class GraphQLMutation implements GraphQLMutationResolver {


    public static TaskDO createTask(TaskDO task) {
        System.out.println("task = " + task);
        return TaskService.addOne(task);
    }

    public static Boolean deleteTask(String sid) {
        System.out.println("sid = " + sid);
        return TaskService.deleteOne(sid);
    }

}
