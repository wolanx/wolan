package com.zx5435.wolan.graph;

import com.coxautodev.graphql.tools.GraphQLMutationResolver;
import com.zx5435.wolan.model.TaskDO;
import org.springframework.stereotype.Component;

import java.io.IOException;


@Component
public class GraphQLMutation implements GraphQLMutationResolver {


    public static TaskDO createTask(TaskDO task) throws IOException {
        System.out.println("task = " + task);
        return GraphQLQuery.getTaskBySid("01-fs");
    }

}
