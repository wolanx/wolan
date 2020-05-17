package com.zx5435.wolan.git;

import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;
import org.eclipse.jgit.api.CloneCommand;
import org.eclipse.jgit.api.Git;

import java.io.File;

public class GitService {

    public static boolean taskDoClone(TaskDO task) {
        String url = task.getGit().getUrl(); // https://github.com/zx5435/go-fs.git
        if (url.isEmpty()) {
            return false;
        }
        GitService.cloneRepository(url, WoConf.WorkPath + "/" + task.getSid() + "/code");
        return true;
    }

    private static void cloneRepository(String url, String localPath) {
        try {
            CloneCommand cc = Git.cloneRepository().setURI(url);
            cc.setDirectory(new File(localPath)).call();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}
