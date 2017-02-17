import rethinkdb

def handler(conn, event):
    try:
        rethinkdb.db("test").table_create("test_table").run(conn)
        return rethinkdb.table("test_table").insert([{ "fieldA": 1234 }]).run(conn)
    except Exception as e:
        return {'error': str(e)} 
