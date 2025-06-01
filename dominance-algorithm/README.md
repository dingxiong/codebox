## A Simple, Fast Dominance Algorithm

Page 5 gives the pseudocode of the classic iterative algorithm to compute dominators. However, I was confused by the 
initialization. Then chatgpt shows me a more detailed version
```
for all nodes n:
    DOM[n] ← {all nodes}   // Start pessimistically

DOM[start] ← {start}       // Start dominates itself only

Changed ← true
while Changed:
    Changed ← false
    for all nodes n ≠ start, in reverse post-order:
        new_DOM ← ∩ (DOM[p]) for all p ∈ preds(n)
        new_DOM ← new_DOM ∪ {n}
        if new_DOM ≠ DOM[n]:
            DOM[n] ← new_DOM
            Changed ← true
```

Ah! The start node has index 0, so the initialization from 1 to N in the pager 
makes sense then.
