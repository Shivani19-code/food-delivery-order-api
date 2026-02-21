# Prompts Log: Food Delivery Order API

This document records the interactions used to generate this codebase using AI assistance.

## Prompt 1: Project Initialization & Analysis
**Prompt**: "Dear Students... I have chosen food-delivery-order-api. please make it attractive, interviewer accepting, expecting needs of interviewer."
**Action**: Analyzed requirements, searched for industry standards, and proposed a Clean Architecture implementation plan.

## Prompt 2: Core Structure & Authentication
**Prompt**: "How to run this" (implied "Proceed with implementation")
**Action**:
- Initialized Go module and directory structure.
- Implemented JWT authentication and middleware.
- Created User models and services.
- Set up GORM with SQLite and auto-migrations.

## Prompt 3: Business Features
**Action**:
- Implemented Restaurant and Menu CRUD operations.
- Implemented Order management with transactions.
- Created seed data function for immediate reviewer testing.

## Prompt 4: Documentation & Finalization
**Action**:
- Created `README.md` with detailed usage instructions.
- Created `design_doc.md` explaining architectural decisions.
- Organized folder structure for final submission.
- Ran `go mod tidy` for dependency resolution.
