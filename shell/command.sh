export out="--dry-run=client -o yaml"
kubectl create deploy redis-ds --image=redis:5-alpine $out