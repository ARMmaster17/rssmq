from watergrid.pipelines import StandalonePipeline
from watergrid.pipelines.pipeline import Pipeline

from app.steps.get_sources_step import GetSourcesStep


class App:
    def __init__(self):
        pass

    def build_pipeline(self) -> Pipeline:
        pipeline = StandalonePipeline('rssmq_pipeline')
        sources = ["https://xkcd.com/rss.xml"]
        pipeline.add_step(GetSourcesStep(sources))
        return pipeline