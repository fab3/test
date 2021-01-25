from flask import Flask

app = Flask(__name__)
from flask_sqlalchemy import SQLAlchemy

app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///data.db"
db = SQLAlchemy(app)

##link to a db using a model
class Drink(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), unique=True, nullable=False)
    descript = db.Column(db.String(150))

    def __repr__(self):
        return f"{self.name} - {self.descript}"


@app.route("/")
def index():
    return "HI there i am new in here INDEX !"


@app.route("/drinks")
def get_drinks():
    drink = Drink.query.all()
    result = []
    for drink in drink:
        result_data = {"name": drink.name, "descr": drink.descript}
        result.append(result_data)
    return {"drinks": result}


@app.route("/drinks/<id>")
def get_id(id):
    drink = Drink.query.get_or_404(id)
    return {"name": drink.name}
