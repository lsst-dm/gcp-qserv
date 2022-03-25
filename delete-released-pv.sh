kubectl get pv  | grep Released | cut -d' ' -f1 | xargs kubectl delete pv -
