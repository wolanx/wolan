import com.zx5435.wolan.git.GitService;
import com.zx5435.wolan.graph.GraphQLQuery;
import com.zx5435.wolan.model.TaskDO;

import java.io.FileNotFoundException;

public class Git {

    public static void main(String[] args) {
        try {
            TaskDO task = GraphQLQuery.taskGetBySid("fs");
            GitService.taskDoClone(task);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

    }

}
