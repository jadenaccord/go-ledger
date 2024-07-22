# go-ledger

A simple CLI (and web-GUI) ledger application for small businesses.

- Keep your data in CSV or SQLite
- CLI for straightforward and automatable operations
- Web-GUI for easy-to-use interface

## Roadmap

### MVP

CLI (v0.1.0)

- [ ] Create and open databases
- [ ] Create and open tables
- [ ] Print tables
- [ ] Insert entries into tables
- [ ] Print tables, filtered
- [ ] Edit entries
- [ ] Delete entries

GUI (v0.2.0)

- [ ] Display table
- [ ] New invoice form
- [ ] New transaction form

### Full Product

(v1.0.0)

- [ ] All functions in CLI and GUI
- [ ] Generate reports (balance sheet, income statement, cash flow, etc.)
- [ ] Generate reports to PDF
- [ ] Lock table entries
- [ ] (Automatic) DB backups (?)
- [ ] Git integration (?)

## Accounting System

go-ledger uses double-entry accounting based on two cycles: invoicing/billing cycling, and cash cycle. I.e., invoices and bills are entered first, the payment is entered later (although it can be on the same day).

