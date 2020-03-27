package com.zx5435.wolan.k8s;

import io.kubernetes.client.ApiClient;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.AppsV1Api;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.*;
import io.kubernetes.client.util.Config;
import io.kubernetes.client.util.Yaml;

import java.io.File;
import java.io.IOException;

/**
 * https://github.com/kubernetes-client/java
 */
public class K8s {

    private static final String KUBECONFIG = "./src/main/java/com/zx5435/wolan/k8s/kubeconfig.yaml";

    static {
        ApiClient client = null;
        try {
            client = Config.fromConfig(KUBECONFIG);
        } catch (IOException e) {
            e.printStackTrace();
        }
        Configuration.setDefaultApiClient(client);
    }

    public static void main(String[] args) throws ApiException, IOException {
        AppsV1Api api = new AppsV1Api();

        File file = new File("./src/main/java/com/zx5435/wolan/k8s/go-fs.yaml");
        V1Deployment yamlObj = (V1Deployment) Yaml.load(file);

        V1Deployment aDefault = api.createNamespacedDeployment("default", yamlObj, null, null, null);
        System.out.println("aDefault = " + aDefault);
    }

    public static void listNs() throws ApiException {
        CoreV1Api api = new CoreV1Api();

        V1NamespaceList arr = api.listNamespace(null, null, null, null, null, null, null, null, null);

        for (V1Namespace one : arr.getItems()) {
            System.out.println("one = " + one.getMetadata().getName());
        }
    }

    public static void listPod() throws ApiException {
        CoreV1Api api = new CoreV1Api();
        V1PodList arr = api.listPodForAllNamespaces(null, null, null, null, null, null, null, null, null);
        for (V1Pod one : arr.getItems()) {
            System.out.println(one.getMetadata().getName());
        }
    }

}

