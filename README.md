# Sisyphus

A **SIEM**-like event processing and alerting system focused on continuously deriving operational state from immutable signals.

```mermaid
flowchart LR
    subgraph Ingestion
        Source["Event Sources<br/>(hosts, apps, network)"]
        Signal["Signal<br/>(immutable event)"]
        Source --> Signal
    end

    subgraph Processing
        Metric["Metric<br/>(derived state)"]
        Evaluator["Evaluator<br/>(rules / thresholds)"]
        Signal --> Metric
        Metric --> Evaluator
    end

    subgraph Response
        Alarm["Alarm<br/>(stateful alert)"]
        Action["Action<br/>(notify / remediate)"]
        Evaluator --> Alarm
        Alarm --> Action
    end

    subgraph Feedback
        ActionSignal["Action Result<br/>(new signal)"]
        Action --> ActionSignal
        ActionSignal --> Signal
    end

```

## Glossary

- **Signal** - Immutable and temporally decoupled event from a host, app, or network. Signals are append only and serve as a source of truth.
- **Metric** - A derived property computed from one or more Signals, calculated over defined time windows and persisted for historical analysis.
- **Alarm** - A stateful, actionable condition triggered when user-defined rules over Metrics are violated. Alarms track lifecycle state (e.g. firing, resolved).
- **Action** - A response executed as a result of an Alarm, such as notification or automated remediation, which itself emits a new Signal back into the system.