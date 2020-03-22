import com.zx5435.wolan.git.GitService;
import com.zx5435.wolan.graph.TaskGraph;
import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;

import java.io.FileNotFoundException;

public class Git {

    public static void main(String[] args) {
        try {
            TaskDO task = TaskGraph.getTaskByName("fs");
            GitService.taskDoClone(task);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

    }

}
