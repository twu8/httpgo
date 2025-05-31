# Manual Verification Steps for Version Endpoint

This document outlines the steps to manually verify that the version information is correctly injected into the application via Docker build arguments and exposed through the `/version` HTTP endpoint.

1.  **Build the Docker image using `Dockerfile2`:**
    Open a terminal in the root of the project and run the following command. This command builds the Docker image and passes build arguments to set the version, build time, and commit hash.

    ```bash
    docker build -f Dockerfile2 \
      --build-arg VERSION_ARG="v1.2.3" \
      --build-arg BUILD_TIME_ARG="$(date -u +'%Y-%m-%dT%H:%M:%SZ')" \
      --build-arg COMMIT_HASH_ARG="$(git rev-parse --short HEAD)" \
      -t httpgo-versioned .
    ```
    *(Note: The `COMMIT_HASH_ARG` command assumes `git` is installed and the project is a git repository. If not, you can replace `$(git rev-parse --short HEAD)` with a static string like "testcommithash" for testing purposes.)*

2.  **Run the Docker container:**
    After the image is successfully built, run a container from it:
    ```bash
    docker run -d -p 8086:8086 --name httpgo-test httpgo-versioned
    ```
    This command starts the container in detached mode (`-d`) and maps port `8086` of the container to port `8086` on the host.

3.  **Access the `/version` endpoint:**
    Once the container is running, you can access the `/version` endpoint. Open a web browser and navigate to `http://localhost:8086/version`, or use `curl` in your terminal:
    ```bash
    curl http://localhost:8086/version
    ```

4.  **Verify the output:**
    The command should return a JSON response. The `version` should be "v1.2.3", and `build_time` and `commit_hash` should reflect the values passed during the `docker build` command (or the current time/git hash if you used the dynamic commands).

    Example expected output:
    ```json
    {
      "version": "v1.2.3",
      "build_time": "YYYY-MM-DDTHH:MM:SSZ",
      "commit_hash": "actualcommithash"
    }
    ```
    (The `build_time` will be the time of the build, and `commit_hash` will be the short hash of the current HEAD commit if `git` was used.)

5.  **Clean up:**
    After you have verified the endpoint, you should stop and remove the Docker container to free up resources.
    ```bash
    docker stop httpgo-test
    docker rm httpgo-test
    ```
    Optionally, if you no longer need the Docker image, you can remove it:
    ```bash
    docker rmi httpgo-versioned
    ```
