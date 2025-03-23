# timepad-queue-manager
Online queue manager for Timepad events

Checklist:
0. Prepare the infrastructure: set up a database (tables ar visitors and list of tables people queue up to), Telegram bot account, set up an application itself.
1. Read the SQL database file, which was formed from spreadsheet (downloadable from Timepad website).
2. Create a simple Telegram bot and check if it has access to list of users. Use their names OR other unique identifier to prevent abuse.
3. Add a way to render list of tables.
4. Implement queue itself.
5. Test the app. The *user* should be able to:
    5.1. Login via it's unique identifier.
    5.2. Get a list of available tables.
    5.3. Book a table for 45 minutes.
    5.4. Remove a booking before the start of the session.
    5.5. Set a state, as if the *user* most likely would visit a session, not sure yet or would probably miss the session.
6. Test the app. The *admin* should be able to:
    6.1. Login via it's unique identifier and password.
    6.2. Get a list of sessions.
    6.3. Get a list of current bookings for a single session and all sessions (not at the same time).
    6.4. Add a guest on a session even if the limit is reached.
    6.5. Manage user states: attended a session, missed a session, uncertain.
7. Deploy the app.