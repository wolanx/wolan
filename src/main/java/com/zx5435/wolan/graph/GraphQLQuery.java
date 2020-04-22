package com.zx5435.wolan.graph;

import com.coxautodev.graphql.tools.GraphQLQueryResolver;
import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;
import com.zx5435.wolan.service.TaskService;
import org.springframework.stereotype.Component;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;


/**
 * # GraphQL从入门到放弃 https://xeblog.cn/articles/6
 * # GraphQL 零基础教程 https://segmentfault.com/a/1190000021899271
 */
@Component
public class GraphQLQuery implements GraphQLQueryResolver {

    public static TaskDO taskGetBySid(String sid) {
        System.out.println("sid = " + sid);
        return TaskService.getOne(sid);
    }

    public List<TaskDO> taskList() {
        ArrayList<TaskDO> ret = new ArrayList<>();

        File workFile = new File(WoConf.WorkPath);
        File[] taskFiles = workFile.listFiles();

        for (File taskFile : Objects.requireNonNull(taskFiles)) {
            if (taskFile.isDirectory()) {
                ret.add(taskGetBySid(taskFile.getName()));
            }
        }

        return ret;
    }

}
