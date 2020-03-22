package com.zx5435.wolan.git;

import com.zx5435.wolan.model.TaskDO;
import com.zx5435.wolan.other.WoConf;
import org.eclipse.jgit.api.CloneCommand;
import org.eclipse.jgit.api.Git;

import java.io.File;

public class GitService {

    public static void taskDoClone(TaskDO task) {
        String url = task.getGit().getUrl(); // https://github.com/zx5435/go-fs.git
        GitService.cloneRepository(url, WoConf.WorkPath + "/" + task.getSid() + "/code");
    }

    public static String cloneRepository(String url, String localPath) {
        try {
            System.out.println("开始下载......");

            CloneCommand cc = Git.cloneRepository().setURI(url);
            cc.setDirectory(new File(localPath)).call();

            System.out.println("下载完成......");

            return "success";
        } catch (Exception e) {
            e.printStackTrace();
            return "error";
        }
    }

}
