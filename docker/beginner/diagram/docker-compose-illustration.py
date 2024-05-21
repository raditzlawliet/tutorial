# diagram-compose.py
from diagrams import Cluster, Diagram, Edge
from diagrams.aws.compute import ECS
from diagrams.aws.database import RDS
from diagrams.aws.network import ELB
from diagrams.onprem.container import Docker
from diagrams.onprem.database import Postgresql
from diagrams.programming.language import Go
from diagrams.generic.blank import Blank
from diagrams.generic.storage import Storage
from diagrams.onprem.client import User
from diagrams.programming.flowchart import Document

with Diagram("Docker Compose illustration", show=True):
    user = User("user")

    with Cluster("Infrastructure"):
        volume = Storage("Persistent Volume (pg_data)")
        dockerfile = Document("app Dockerfile")
        c_infra = [volume, dockerfile]

        with Cluster(""):
            app = Docker("App (app)")
            database = Docker("Database (postgresdb)")
            c_svc = [app, database]

    dockerfile >> Edge(label="build image", color="red") >> app
    (
        user
        >> Edge(label="http")
        >> app
        >> Edge(label="network=learning")
        >> database
        >> Edge(label="read + write")
        >> volume
    )
