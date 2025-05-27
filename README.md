# go-masters-2025

## Задание к занятию #10
- Развернуть на своем компьютере Kubernetes (Kind, Minikube: https://kubernetes.io/docs/tasks/tools/).
- Собрать образ для демонстрационного приложения.
- Развернуть приложение из образа в K8s с помощью манифеста в виде deployment.
- Увеличить/уменьшить количество реплик приложения



### Сборка и загрузка

- ``docker build -t localhost:5000/go-masters:1.0.0 .``
- ``docker push localhost:5000/go-masters:1.0.0``        # Kind
- ``minikube image load localhost:5000/go-masters:1.0.0``  # для Minikube


### Развёртывание и масштабирование

- ``kubectl apply -f deploy/namespace.yaml``
- ``kubectl apply -f deploy/``


- ``kubectl -n go-masters get deploy,pods,svc``
- ``kubectl -n go-masters scale deploy go-masters --replicas=5``


#### Добавлено:
- Namespace go-masters для изоляции ресурсов и RBAC
- Immutable tag 1.0.0 -	воспроизводимые deploy’и
- Multistage + distroless	минимальный attack-surface
- readiness/liveness probes	zero-downtime rollout и auto-heal
- revisionHistoryLimit - быстрый rollback
- HPA -	автомасштаб по CPU, production-friendly
- Prometheus annotations - мониторинг без лишних конфигов
- Graceful shutdown в Go - не рвём соединения при SIGTERM