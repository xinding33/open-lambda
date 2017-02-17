import json
from bson import json_util


def handler(conn, event):
    try:
        db = conn['test_db']
        collection = db['test_col']
        doc1 = {"field_a": 1234}
        doc1_id = collection.insert(doc1)
        # ObjectId needs a special codec
        # json.dumps(doc1_id, default=json_util.default)
        return str(doc1_id)
    except Exception as e:
        return {'error': str(e)} 
