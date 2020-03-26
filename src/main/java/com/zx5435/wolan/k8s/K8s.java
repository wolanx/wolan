package com.zx5435.wolan.k8s;

import io.kubernetes.client.ApiClient;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1Namespace;
import io.kubernetes.client.models.V1NamespaceList;
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
        ApiClient client = null;
        try {
            client = Config.fromConfig(KUBECONFIG);
        } catch (IOException e) {
            e.printStackTrace();
        }
        Configuration.setDefaultApiClient(client);
    }

    public static void main(String[] args) throws ApiException {
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

