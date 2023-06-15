oc create -f config/wmco/wmco-namespace.yaml
oc create -f config/wmco/wmco-og.yaml
oc create -f config/wmco/wmco-sub.yaml

# be sure that this ssh key is an RSA private key
oc create secret generic cloud-private-key --from-file=private-key.pem=${HOME}/.ssh/openshift.pem \
    -n openshift-windows-machine-config-operator 
