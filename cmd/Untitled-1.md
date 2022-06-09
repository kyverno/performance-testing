# https://github.com/kyverno/kyverno/issues/4072

# reproduce the issue
Items:
- policies: https://github.com/kyverno/kyverno/files/8838363/policies.yaml.txt [X]
- duration: 1 hour [X]
- scales: xl (3k total workloads) / 600 cronjob [X]
- exclude cronjob from steps down process [X]
- version: 1.7.0 [X]

Pendings:
- implementation PR to kyverno/kyverno
- KDP to kyverno/kdp
- doc updates to kyverno/website