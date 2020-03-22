package com.zx5435.wolan.model;


import lombok.Data;

@Data
public class TaskDO {

    /**
     * version: "1.0"
     * name: "go-fs"
     * git:
     * url: 'https://github.com/zx5435/go-fs.git'
     * branch: 'master'
     * docker-compose: "docker-composer.yml"
     */

    private String sid;
    private String version;
    private String name;

    private GitBean git;

    @Data
    public static class GitBean {
        private String url;
        private String branch;
    }

}
