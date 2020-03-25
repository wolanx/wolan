package com.zx5435.wolan.k8s;

import io.kubernetes.client.ApiClient;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1Pod;
import io.kubernetes.client.models.V1PodList;
import io.kubernetes.client.util.Config;

import java.io.IOException;

/**
 * https://github.com/kubernetes-client/java
 */
public class K8s {

    private static final String KUBECONFIG = "./src/main/java/com/zx5435/wolan/k8s/kubeconfig.yaml";

    static {
        ApiClient client = Config.fromConfig(KUBECONFIG);
        Configuration.setDefaultApiClient(client);
    }

    public static void main(String[] args) throws IOException, ApiException {

    }

    public static void listPod() throws IOException, ApiException {

        CoreV1Api api = new CoreV1Api();
        V1PodList list = api.listPodForAllNamespaces(null, null, null, null, null, null, null, null, null);
        for (V1Pod item : list.getItems()) {
            System.out.println(item.getMetadata().getName());
        }
    }

}

