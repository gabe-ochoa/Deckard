# Deckard - a replicate killer

Kill kube pods automagically.

A small Go tool to restart deployments, daemonsets, or pods en mass.

The tool respects the restartStrategy assigned to the deployment. If no strategy is defined, a default of 1 up/1 down is used.

Use cases:
- Restarting on a cadance for a clean slate for stateless applications
- Restarting an application without having to trigger a change in the Controller Manager
- Forcing a pod to pick up configmap changes

Syntax:

`deckard target deployment my-app`
`deckard target daemonset my-app`

