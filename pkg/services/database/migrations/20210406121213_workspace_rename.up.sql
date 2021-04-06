BEGIN;

ALTER TABLE companies RENAME COLUMN company_id TO workspace_id;

ALTER TABLE companies RENAME TO workspaces ;

ALTER TABLE account_company RENAME COLUMN company_id TO workspace_id;

ALTER TABLE account_company RENAME TO account_workspace ;

COMMIT;
